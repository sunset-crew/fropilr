Name:           fropilr
Version:        1
Release:        v0.1.2%{?dist}
Summary:        A profile management system

Group:          Simple
License:        GPL
URL:            https://github.com/kf4jas/fropilr
Source0:        fropilr-1.tar.gz

BuildArch:      noarch
#BuildRequires: systemd-rpm-macros
#%{?systemd_requires}

# Requires: python-inotify

%description
A simple method of profile management

%prep
%setup -q

%build
# %%configure
# make %{?_smp_mflags}
make build

%install
install -m 0755 -d $RPM_BUILD_ROOT/usr/local/bin
install -m 0755 -d $RPM_BUILD_ROOT/usr/local/fropilr
install -m 0755 fropilr $RPM_BUILD_ROOT/usr/local/bin/fropilr

%clean
rm -rf $RPM_BUILD_ROOT

%files
%defattr(-,root,root,-)
%attr(-,root,root) /usr/local/bin/fropilr

%pre
if [ $1 == 2 ]; then
  # this only runs during an update
  echo "updating"
fi

%post
echo %{release} > /usr/local/fropilr/version

%preun
echo "prerun"

%postun
if [  $1 == 1 ]; then
  echo "update 2"
fi

%changelog
